#pragma once

#include "../tools/definitions.hpp"
#include "app.hpp"

namespace Vuelto {

class Renderer2D;

namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win);

}  // namespace Application

class Renderer2D {
 public:
  void DrawRect(float x, float y, float height, float width, float color1, float color2, float color3);
  void SetBackgroundColor(float color1, float color2, float color3);
  void DrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3);
};

}  // namespace Vuelto
