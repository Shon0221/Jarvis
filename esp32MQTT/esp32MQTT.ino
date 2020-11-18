#include <WiFi.h>
#include <PubSubClient.h>

const int RELAY = 15;
//SSID
const char* ssid = "";
//SSID PW
const char* password = "";
//MQTT Host 
const char* mqttServer = "104.199.185.7";
/ MQTT Port
const int mqttPort = 1883;
//MQTT username
const char* mqttUser = "iot";
//MQTT PW
const char* mqttPassword = "iot@1qaz2wsx";
// Relay 狀態預設為 close
char* RELAYStatus = "close";
// Relay TOPIC
char* RelayTopic = "IOT/Relay/Switch01/Controlle";

WiFiClient espClient;
PubSubClient client(espClient); {

  void callback(char* topic, byte * payload, unsigned int length) {
    String payloadStr;
    Serial.print("目前 TOPIC : ");
    Serial.println(topic);
    Serial.print("接收到的訊息: ");

    for (int i = 0; i < length; i++) {
      payloadStr = payloadStr + String((char)payload[i]);
    }

    Serial.println(payloadStr);
    Serial.println("-------------------------");
    if (topic == RelayTopic ) {
      if (payloadStr == "close") {
        RELAYStatus = "close";
        digitalWrite(RELAY, HIGH);
      }
      else if (payloadStr == "open") {
        RELAYStatus = "open";
        digitalWrite(RELAY, LOW);
      }
      //送出目前的狀態
      //client.publish("IOT/Relay/Switch01/Status", RELAYStatus);
    }
  }

  void setup() {
    //設定鮑率
    Serial.begin(115200);
    //設定腳位為輸出模式
    pinMode(RELAY, OUTPUT);
    //初始輸出為 HIGH , 這樣讓 ESP32 啟動時 , 讓 Relay 在常閉狀態
    digitalWrite(RELAY, HIGH);

    //Wifi 開始連線
    WiFi.begin(ssid, password);
    //Wifi 連線結果沒有成功一直輸出下列文字
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.println("跟 Wifi 連線中...");
    }
    //Wifi 連線成功 
    Serial.println("Wifi 連線成功");

    //MQTT 連線開啟
    client.setServer(mqttServer, mqttPort);
    //MQTT Callback
    client.setCallback(callback);
    
    //MQTT 連線直到成功為止
    while (!client.connected()) {
      //開始連線 
      Serial.println("跟 MQTT Server 連線中...");
      //輸入 MQTT 帳號密碼
      if (client.connect("ESP32Client", mqttUser, mqttPassword)) {       
        //MQTT 連線成功
        Serial.println("MQTT 連線成功");
        delay(500);
      } else {        
        //印出連線失敗原因
        Serial.print("failed with state: ");
        Serial.print(client.state());
      }
    }
  }

  void loop() {
    client.loop();
  }
