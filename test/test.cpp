#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", true);

  Vuelto::Renderer2D renderer = Vuelto::Application::CreateRenderer2D(win);

  while (!win.WindowShouldClose()) {
    renderer.DrawRect(0.1, 0.1, 0.1, 0.1, 0.1, 0.1, 0.1);
    win.Refresh();
  }

  Vuelto::Application::Terminate();
  return 0;
}
