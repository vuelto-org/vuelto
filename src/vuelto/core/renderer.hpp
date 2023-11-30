#pragma once

#include "../tools/definitions.hpp"

namespace Vuelto {
class Renderer2D {
 public:
  void DrawRect(float x, float y, float height, float width, float color1, float color2, float color3);
  void SetBackgroundColor(float color1, float color2, float color3);
  void DrawLine(int x1, int x2, int y1, int y2, float color1, float color2, float color3);
};

class SoftwareRenderer {
  // Init function for the Software renderer (Keep in mind that the software renderer does not support window resizing,
  // and it will be automatically turned off. This is to prevent weird bugs or issues with the renderer or code)
 public:
  void Init(int width, int height);
  void SetPixel(int x, int y, unsigned char r, unsigned char g, unsigned char b);
  void DrawRect(int x, int y, int width, int height, int r, int g, int b);
  void SetBackgroundColor(unsigned char r, unsigned char g, unsigned char b);
  void Refresh();
  void Terminate();
  void ResizeBuffer(int newWidth, int newHeight);
  bool Initialized();
};

}  // namespace Vuelto
