#include "../src/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", true);

  Vuelto::SoftwareRenderer renderer = Vuelto::Application::CreateSoftwareRenderer(win);

  while (!win.WindowShouldClose()) {
    renderer.DrawRect(100, 100, 200, 200, 123, 123, 132);
    win.Refresh();
  }

  Vuelto::Application::Terminate();
  return 0;
}
