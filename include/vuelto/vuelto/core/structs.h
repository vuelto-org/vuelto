#pragma once

#include <GLFW/glfw3.h>

typedef struct {
  GLFWwindow *window;
  int width;
  int height;
  const char *title;
} Vuelto_Window;

typedef struct {
  float width, height, x, y;
  GLuint texture;
} Vuelto_Image;
