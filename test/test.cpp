#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", true);

  while (!win.WindowShouldClose()) {
    Vuelto::Renderer2D::DrawRect(-0.1, -0.1, 0.5, 0.5, 0.1, 0.4, 0.7);
    win.Refresh();
  }

  Vuelto::Application::Terminate();
  return 0;
}
