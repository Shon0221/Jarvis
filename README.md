# Jarvis

## 麥克風的使用

使用 SpeechRecognizer 訪問麥克風則必須安裝 PyAudio 軟體包，請關閉當前的解釋器窗口，進行以下操作：

### PyAudio

安裝 PyAudio 的過程會因作業系統而異。

#### Debian Linux

Debian的Linux（如 Ubuntu ），則可使用 apt 安裝 PyAudio：

    udo apt-get install python-pyaudio python3-pyaudio

安裝完成後可能仍需要啟用 pip install pyaudio ，尤其是在虛擬情況下運行

#### macOS

macOS 用戶則首先需要使用 Homebrew 來安裝 PortAudio，然後調用 pip 命令來安裝 PyAudio

    brew install portaudio
    pip install pyaudio

#### Windows

Windows 用戶可直接調用 pip 來安裝 PyAudio。

    pip install pyaudio

安裝測試

安裝了 PyAudio 後可從控制台進行安裝測試。

    python -m speech_recognition

確保默認麥克風打開並取消靜音，若安裝正常則應該看到如下所示的內容：

    A moment of silence, please...
    Set minimum energy threshold to 600.4452854381937
    Say something!

對著麥克風講話並觀察 SpeechRecognition 如何轉錄你的講話

### Microphone

    import speech_recognition as sr
    r = sr.Recognizer()
    # 創建一個Microphone 類的實例來訪問它
    mic = sr.Microphone()
    # 取麥克風名稱列表
    print(sr.Microphone.list_microphone_names())

EX:

['HDA Intel PCH: ALC272 Analog (hw:0,0)', 'HDA Intel PCH: HDMI 0 (hw:0,3)', 'sysdefault', 'front', 'surround40', 'surround51', 'surround71', 'hdmi', 'pulse', 'dmix', 'default']

list_microphone_names（）返回列表中麥克風設備名稱的索引。在上面的輸出中，如果要使用名為 「front」 的麥克風，該麥克風在列表中索引為 3，則可以創建如下所示的麥克風實例：

    mic = sr.Microphone(device_index=3)

#### 使用 listen（）獲取麥克風輸入數據

準備好麥克風實例後，讀者可以捕獲一些輸入。

就像 AudioFile 類一樣，Microphone 是一個上下文管理器。可以使用 with 塊中 Recognizer 類的 listen（）方法捕獲麥克風的輸入。該方法將音頻源作為第一個參數，並自動記錄來自源的輸入，直到檢測到靜音時自動停止

    with mic as source:
      audio = r.listen(source)

執行 with 塊後請嘗試在麥克風中說出 「hello」 。請等待解釋器再次顯示提示，提示返回就可以識別語音

    r.recognize_google(audio)

要處理環境噪聲，可調用 Recognizer 類的 adjust_for_ambient_noise（）函數，其操作與處理噪音音頻文件時一樣。由於麥克風輸入聲音的可預測性不如音頻文件，因此任何時間聽麥克風輸入時都可以使用此過程進行處理

    with mic as source:
      r.adjust_for_ambient_noise(source)
      audio = r.listen(source)

    
