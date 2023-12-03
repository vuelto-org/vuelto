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
  void DrawRectWithShader(float x, float y, float width, float height, const char* colorShader);
  void SetBackgroundColor(float color1, float color2, float color3);

  Image LoadImage(const char* imagePath, int x, int y, int& width, int& height);
};

class Renderer2D::Image {
 public:
  float width, height, x, y;
  unsigned int texture;

  void DrawImage();

};

}  // namespace Vuelto
