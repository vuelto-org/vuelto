# Framerate in Vuelto
Hello and welcome to another Developer part of vuelto! In this one we are going to quickly explain how we manage the framerate in vuelto and how the end users experience it!

## Introduction
First up, we have to get some things straight. For framerate, we need the following parts:
1. User-set settings
1. Calculate correct sleep timing
3. Freeze main loop to keep framerate steady

### Number one
Example is SetFPS. This is just a couple of functions to manage the framerate inside of vuelto. We receive the amount of frames per second that you want, and we time the freezing (point number 3) correctly so that there is a steady framerate.

### Number two
Calculating the required time to freeze the process in order to keep that framerate. This has to be done accurately, else you can result in stutters.

### Number three
Last part is that we freeze the main loop in order to stay at the requested framerate. This should **not** be done when the current framerate is already lower than requested. In that case, the user/dev will be warned, with the possibility to silent the warnings.

