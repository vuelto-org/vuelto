#pragma once

#include "app.hpp"

namespace Vuelto {

class Renderer2D;

namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win);

}  // namespace Application

class Renderer2D {
 public:
  void DrawRect(float x, float y, float width, float height, float r, float g, float b);
  void DrawRectWithShader(float x, float y, float width, float height, const char* colorShader);
  void SetBackgroundColor(float color1, float color2, float color3);
};

}  // namespace Vuelto
