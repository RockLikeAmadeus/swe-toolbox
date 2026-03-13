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

# Reading Analog Signals

# Debugging with the Serial Monitor

When connected to our machine via USB, the Arduino board internally makes use of a serial-to-USB converter, since the serial interface is what the MCU uses for more sophisticated messaging.