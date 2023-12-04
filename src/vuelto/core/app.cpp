#include "app.hpp"

#include <iostream>

#include "../tools/definitions.hpp"

namespace Vuelto {
namespace Application {

bool MultipleWindowsEnabled = false;

void Init() {
  if (!glfwInit()) {
    std::cout << "GLFW Init failed\n";
  }
}

void InitMultipleWindows() {
  Init();
  MultipleWindowsEnabled = true;
}

void framebuffer_size_callback(GLFWwindow *window, int newWidth, int newHeight) {
  glViewport(0, 0, newWidth, newHeight);
}

Window CreateWindow(int width, int height, const char *title, bool resizable) {
  glfwWindowHint(GLFW_RESIZABLE, resizable);

  GLFWwindow *glfw_window = glfwCreateWindow(width, height, title, NULL, NULL);
  if (!glfw_window) {
    glfwTerminate();
    std::cout << "GLFW Window creation failed\n";
  }

  glfwSetFramebufferSizeCallback(glfw_window, framebuffer_size_callback);

  if (!MultipleWindowsEnabled) glfwMakeContextCurrent(glfw_window);

  Window window;
  window.width = width;
  window.height = height;
  window.title = title;
  window.window = glfw_window;

  return window;
}

void DestroyWindow(Window win) { glfwDestroyWindow(win.window); }

void Terminate() { glfwTerminate(); }

}  // namespace Application

bool Window::WindowShouldClose() {
  while (!glfwWindowShouldClose(window)) {
    return false;
  }
  glfwDestroyWindow(window);
  return true;
}

void Window::MakeContextCurrent() { glfwMakeContextCurrent(window); }

void Window::Refresh() {
  glfwSwapBuffers(window);
  if (!Application::MultipleWindowsEnabled) glfwPollEvents();
  glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
}

}  // namespace Vuelto
