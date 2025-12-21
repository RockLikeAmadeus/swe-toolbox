For quick syntax reference, open File >> Examples in the Arduino IDE.
Also see [this reference](https://docs.arduino.cc/language-reference/).

# Syntax

```arduino
#include "Arduino.h" // Arduino.h is actually included automatically

#define CABIN_LIGHTS_PIN 12

void setup() {
  // put your setup code here, to run once:
  pinMode(CABIN_LIGHTS_PIN, OUTPUT);
}

void loop() {
  // put your main code here, to run repeatedly:
  digitalWrite(CABIN_LIGHTS_PIN, HIGH);
  delay(1000);
  digitalWrite(CABIN_LIGHTS_PIN, LOW);
  delay(100);
}

```
