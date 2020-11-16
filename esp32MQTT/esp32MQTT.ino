#include <WiFi.h>
#include <PubSubClient.h>

const int LED = 15;
const char* ssid = "FengChi Wifi";
const char* password = "70supergod";
const char* mqttServer = "104.199.185.7";
const int mqttPort = 1883;
const char* mqttUser = "iot";
const char* mqttPassword = "iot@1qaz2wsx";

char* LEDStatus = "0";

WiFiClient espClient;
PubSubClient client(espClient);

void callback(char* topic, byte* payload, unsigned int length) {
  String payloadStr;
  Serial.print("Message arrived in topic: ");
  Serial.println(topic);

  Serial.print("Message: ");

  for (int i = 0; i < length; i++) {
    payloadStr = payloadStr + String((char)payload[i]);
  }
  Serial.println(payloadStr);
  Serial.println();
  Serial.println("-------------------------");

  if (payloadStr == "1") {
    LEDStatus = "1";
    digitalWrite(LED, HIGH);
  }
  else {
    LEDStatus = "0";
    digitalWrite(LED, LOW);
  }
  client.publish("IOT/Relay/Switch01/Status", LEDStatus);
}

void setup() {
  Serial.begin(115200);

  pinMode(LED, OUTPUT);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.println("Connect to Wifi...");
  }

  Serial.println("Connected to the Wifi network");
  client.setServer(mqttServer, mqttPort);
  client.setCallback(callback);

  while (!client.connected()) {
    Serial.println("Connecting to MQTT...");
    if (client.connect("ESP32Client", mqttUser, mqttPassword)) {
      Serial.println("connected");
      client.publish("IOT/Relay/Switch01/Status", LEDStatus);
      delay(500);
    } else {
      Serial.print("failed with state: ");
      Serial.print(client.state());
    }
  }

  client.subscribe("IOT/Relay/Switch01/Controlle");
}

void loop() {
  client.loop();
}
