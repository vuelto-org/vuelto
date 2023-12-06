#include "event.h"

void vueltoPollEvents() { glfwPollEvents(); }

bool vueltoKeyPressed(Vuelto_Window window, int keyCode) { return (glfwGetKey(window.window, keyCode) == GLFW_PRESS); }

bool vueltoKeyPressedOnce(Vuelto_Window window, int key) {
  static int keyStates[GLFW_KEY_LAST] = {GLFW_RELEASE};

  if (glfwGetKey(window.window, key) == GLFW_PRESS && keyStates[key] == GLFW_RELEASE) {
    keyStates[key] = GLFW_PRESS;
    return true;
  } else if (glfwGetKey(window.window, key) == GLFW_RELEASE && keyStates[key] == GLFW_PRESS) {
    keyStates[key] = GLFW_RELEASE;
  }
  return false;
}
