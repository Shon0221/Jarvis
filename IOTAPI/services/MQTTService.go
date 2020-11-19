package services

import (
	"IOTAPI/config"
	"IOTAPI/models"
	"fmt"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Publish :
func Publish(mqtt *models.MQTT) error {
	c, err := connectMQTT()
	if err != nil {
		return err
	}
	// 發佈消息
	token := c.Publish(mqtt.Topic, 0, false, mqtt.Message)
	token.Wait()

	disconnectMQTT(c)

	return nil
}

// Subscribe :
func Subscribe(mqtt *models.MQTT) error {
	c, err := connectMQTT()
	if err != nil {
		return err
	}
	// 發佈消息
	token := c.Publish(mqtt.Topic, 0, false, mqtt.Message)
	token.Wait()

	disconnectMQTT(c)

	return nil
}

//HandlerMessageFunction :
var HandlerMessageFunction mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//fmt.Println("收到訊息開始")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	//return string(msg.Payload())
}

func connectMQTT() (mqtt.Client, error) {
	//mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	//mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	//mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)
	CONFIG := config.New("./config.yaml")
	mqttAddr := "tcp://" + CONFIG.MQTT.Host + ":" + strconv.Itoa(CONFIG.MQTT.Port)
	opts := mqtt.NewClientOptions().AddBroker(mqttAddr).SetClientID("IOTController")
	opts.SetUsername(CONFIG.MQTT.Username).SetPassword(CONFIG.MQTT.Password)
	opts.SetKeepAlive(20 * time.Second)
	//opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}

func disconnectMQTT(c mqtt.Client) {
	// 斷開連接
	c.Disconnect(250)
	//time.Sleep(1 * time.Second)
}

func subscribeMQTT(c mqtt.Client, topic string, qos byte, callback mqtt.MessageHandler) {
	// 訂閱主題
	if token := c.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		//fmt.Println(token.Error())
		os.Exit(1)
		//return token.Error()
	}

	// 取消訂閱
	//defer cancelSubscribeMQTT(c, topic)
}

func cancelSubscribeMQTT(c mqtt.Client, topic string) {
	if token := c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func publishMQTT(c mqtt.Client, topic string, qos byte, msg interface{}) {
	// 發佈消息
	token := c.Publish(topic, 0, false, msg)
	token.Wait()
}
