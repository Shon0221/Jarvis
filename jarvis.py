import pyttsx3 # pip install pyttsx3
import speech_recognition as sr # pip install SpeechRecognition
import wikipedia # pip install wikipedia
import pyautogui # pip install pyautogui
import psutil # pip install psutil
import pyjokes # pip install pyjokes
import datetime
import smtplib
import webbrowser as wb
import os

engine = pyttsx3.init()

def speak(txt):
    engine.say(txt)
    engine.runAndWait()

def now():
    n = datetime.datetime.now().strftime("%Y-%m-%d %I:%M:%S")
    speak(n)

def wishMe():
    speak("Welcome back sir!")
    speak("The current date is ")
    now()
    h = datetime.datetime.now().hour
    if h >= 6 and h < 12:
        speak("Good morning")
    elif h >= 12 and h < 18:
        speak("Good afternoon")
    elif h >= 18 and h <= 24:
        speak("Good evening")
    else:
        speak("Good night") 
    speak("Friday at your service. How I can help you?")

def takeCmd():
    r = sr.Recognizer()
    mic = sr.Microphone()
    with mic as src:
        print("Listening...")
        r.pause_threshold = 1
        r.adjust_for_ambient_noise(src)
        audio = r.listen(src)

    try:
        print("Recognizing...")
        q = r.recognize_google(audio)
        print(q)
    except Exception as e:
        print(e)
        speak("Say that again...")
        return "None"
    
    return q

def searWiki(txt):
    speak("Searching...")
    txt = txt.replace("wikipedia", "")
    r = wikipedia.summary(q, sentences=2)
    speak(r)

def sendEmail(to, content):
    accEmail = ''
    pwdEmail = ''
    try:
        server = smtplib.SMTP('smtp.gamil.com', 587)
        server.ehlo()
        server.login(accEmail, pwdEmail)
        server.sendmail(accEmail, to, content)
        server.close()
        speak("Email sent successfully")
    except Exception as e:
        speak(e)
        speak("Unable to send the message")

def searChrome():
    print("search chrome")
    speak("What should I search?")
    chromePath = 'C:\Program Files (x86)\Google\Chrome\Application\chrome.exe %s'
    ## search keyword
    kw = takeCmd().lower()
    print("search chrome key word [" + kw + "]")
    url = "https://www.google.com.tr/search?q={}".format(kw)
    wb.get(chromePath).open_new_tab(url)

def screenshot():
    img = pyautogui.screenshot()
    img.save('') # <- save file

def cpu():
    usage = str(psutil.cpu_percent())
    speak("CPU is at " + usage)

def jokes():
    speak(pyjokes.get_joke())

if __name__ == "__main__":
    wishMe()

    while True:

        q = takeCmd().lower()
        if "now" in q or "time" in q:
            now()
        elif "wikipedia" in q:
           searWiki(q) 
        # elif "send email" in q:
        #     speak("What should I say?")
        #     content = takeCmd()
        #     to = ''
        #     sendEmail(to, content)
        elif "google" in q:
            searChrome()
        # elif "play songs" in q:
        #     songsDir = ""
        #     songs = os.listdir(songsDir)
        #     os.startfile(os.path.join(songsDir, songs[0]))
        # elif "screenshot" in q:
        #     screenshot()
        #     speak("Screenshot done")
        elif "cpu" in q:
            cpu()
        elif "joke" in q:
            jokes()
        elif "offline" in q:
            quit()