#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(500, 500, "test");

  while (!win.WindowShouldClose()) {
    Vuelto::Renderer2D::DrawRect(0, 0, 0.5, 0.5, 0.5, 0.5, 0.5);
  }

  Vuelto::Application::DestroyWindow(win);
  Vuelto::Terminate();

  return 0;
}
