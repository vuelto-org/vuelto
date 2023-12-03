#include "event.hpp"
#include "../core/app.hpp"

namespace Vuelto {
namespace Events {

void PollEvents() { glfwPollEvents(); }

}  // namespace Events

namespace Input {

bool IsKeyPressed(Vuelto::Window window, int keyCode) {
  return (glfwGetKey(window.window, keyCode) == GLFW_PRESS);
}

bool IsKeyPressedOnce(Vuelto::Window window, int key) {
    static int keyStates[GLFW_KEY_LAST] = {GLFW_RELEASE};

    if (glfwGetKey(window.window, key) == GLFW_PRESS && keyStates[key] == GLFW_RELEASE) {
        keyStates[key] = GLFW_PRESS;
        return true;
    } else if (glfwGetKey(window.window, key) == GLFW_RELEASE && keyStates[key] == GLFW_PRESS) {
        keyStates[key] = GLFW_RELEASE;
    }
    return false;
}

}  // namespace Keyboard

}  // namespace Vuelto
