#pragma once

#include "app.hpp"

namespace Vuelto {

class Renderer2D;

namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win);

}  // namespace Application

class Renderer2D {
 public:
  class Image;

  void DrawRect(float x, float y, float width, float height, float r, float g, float b);
  void SetBackgroundColor(float color1, float color2, float color3);
  void DrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3);

  Image LoadImage(const char* imagePath, float x, float y, float width, float height);
  void CleanUp();
};

class Renderer2D::Image {
 public:
  float width, height, x, y;
  GLuint texture;

  void DrawImage();
};

}  // namespace Vuelto
