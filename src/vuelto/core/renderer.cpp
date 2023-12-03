#include "renderer.hpp"

#include <fstream>
#include <iostream>
#include <sstream>

#include "../shaders/shader.hpp"
#include "../tools/definitions.hpp"
#include "app.hpp"

#define STB_IMAGE_IMPLEMENTATION
#include "../vendor/stb_image/stb_image.h"


namespace Vuelto {
namespace Application {

Vuelto::Renderer2D CreateRenderer2D(Window win) {
  Vuelto::Renderer2D renderer;
  return renderer;
}

}  // namespace Application

const char* vertexShaderSourceRect = Vuelto::Shader::ReadShaderFile("src/vuelto/shaders/vertexRect.glsl");
const char* fragmentShaderSourceRect = Vuelto::Shader::ReadShaderFile("src/vuelto/shaders/fragmentRect.glsl");

const char* vertexShaderSourceImage = Vuelto::Shader::ReadShaderFile("src/vuelto/shaders/vertexRect.glsl");
const char* fragmentShaderSourceImage = Vuelto::Shader::ReadShaderFile("src/vuelto/shaders/fragmentRect.glsl");

void Renderer2D::DrawRect(float x, float y, float width, float height, float r, float g, float b) {
  float vertices[] = {x, y, 0.0f, x + width, y, 0.0f, x + width, y + height, 0.0f, x, y + height, 0.0f};

  unsigned int VBO, VAO;
  glGenBuffers(1, &VBO);
  glGenVertexArrays(1, &VAO);

  glBindVertexArray(VAO);
  glBindBuffer(GL_ARRAY_BUFFER, VBO);
  glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);

  glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
  glEnableVertexAttribArray(0);

  unsigned int vertexShader, fragmentShader, shaderProgram;
  vertexShader = glCreateShader(GL_VERTEX_SHADER);
  glShaderSource(vertexShader, 1, &vertexShaderSourceRect, NULL);
  glCompileShader(vertexShader);

  fragmentShader = glCreateShader(GL_FRAGMENT_SHADER);
  glShaderSource(fragmentShader, 1, &fragmentShaderSourceRect, NULL);
  glCompileShader(fragmentShader);

  shaderProgram = glCreateProgram();
  glAttachShader(shaderProgram, vertexShader);
  glAttachShader(shaderProgram, fragmentShader);
  glLinkProgram(shaderProgram);
  glUseProgram(shaderProgram);

  glDeleteShader(vertexShader);
  glDeleteShader(fragmentShader);

  int colorLocation = glGetUniformLocation(shaderProgram, "color");
  glUniform3f(colorLocation, r, g, b);

  // Draw the rectangle
  glBindVertexArray(VAO);
  glDrawArrays(GL_TRIANGLE_FAN, 0, 4);

  // Cleanup
  glDeleteVertexArrays(1, &VAO);
  glDeleteBuffers(1, &VBO);
  glDeleteProgram(shaderProgram);
}

void Renderer2D::DrawRectWithShader(float x, float y, float width, float height, const char* colorShader) {
  float vertices[] = {x, y, 0.0f, x + width, y, 0.0f, x + width, y + height, 0.0f, x, y + height, 0.0f};

  unsigned int VBO, VAO;
  glGenBuffers(1, &VBO);
  glGenVertexArrays(1, &VAO);

  glBindVertexArray(VAO);
  glBindBuffer(GL_ARRAY_BUFFER, VBO);
  glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);

  glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);
  glEnableVertexAttribArray(0);

  unsigned int vertexShader, fragmentShader, shaderProgram;
  vertexShader = glCreateShader(GL_VERTEX_SHADER);
  glShaderSource(vertexShader, 1, &vertexShaderSourceRect, NULL);
  glCompileShader(vertexShader);

  fragmentShader = glCreateShader(GL_FRAGMENT_SHADER);
  glShaderSource(fragmentShader, 1, &colorShader, NULL);
  glCompileShader(fragmentShader);

  shaderProgram = glCreateProgram();
  glAttachShader(shaderProgram, vertexShader);
  glAttachShader(shaderProgram, fragmentShader);
  glLinkProgram(shaderProgram);
  glUseProgram(shaderProgram);

  glDeleteShader(vertexShader);
  glDeleteShader(fragmentShader);

  // Draw the rectangle
  glBindVertexArray(VAO);
  glDrawArrays(GL_TRIANGLE_FAN, 0, 4);

  // Cleanup
  glDeleteVertexArrays(1, &VAO);
  glDeleteBuffers(1, &VBO);
  glDeleteProgram(shaderProgram);
}

void Renderer2D::SetBackgroundColor(float color1, float color2, float color3) {
  glClearColor(color1, color2, color3, 1);
}

Renderer2D::Image Renderer2D::LoadImage(const char* imagePath, int x, int y, int& width, int& height) {
      // Load image using stb_image
    int channels;
    unsigned char* image = stbi_load(imagePath, &width, &height, &channels, STBI_rgb_alpha);

    if (!image)
    {
        std::cerr << "Failed to load image: " << imagePath << std::endl;
    }

    // Generate texture
    unsigned int texture;
    glGenTextures(1, &texture);
    glBindTexture(GL_TEXTURE_2D, texture);
    glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0, GL_RGBA, GL_UNSIGNED_BYTE, image);
    glGenerateMipmap(GL_TEXTURE_2D);

    // Free image data
    stbi_image_free(image);
    Renderer2D::Image img;
    img.texture = texture;
    img.x = x;
    img.y = y;
    img.width = width;
    img.height = height;

    return img;
}

void Renderer2D::Image::DrawImage() {
    // Set up vertex data, buffers, etc.
    float vertices[] = {
        // Positions      // Texture Coords
        x, y, 0.0f, 0.0f, 0.0f,
        x + width, y, 0.0f, 1.0f, 0.0f,
        x + width, y + height, 0.0f, 1.0f, 1.0f,
        x, y + height, 0.0f, 0.0f, 1.0f
    };

    unsigned int VBO, VAO;
    glGenVertexArrays(1, &VAO);
    glGenBuffers(1, &VBO);

    glBindVertexArray(VAO);
    glBindBuffer(GL_ARRAY_BUFFER, VBO);
    glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW);

    // Position attribute
    glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 5 * sizeof(float), reinterpret_cast<void*>(0));
    glEnableVertexAttribArray(0);

    // Texture coordinate attribute
    glVertexAttribPointer(1, 2, GL_FLOAT, GL_FALSE, 5 * sizeof(float), reinterpret_cast<void*>(3 * sizeof(float)));
    glEnableVertexAttribArray(1);

    // Compile Shaders
    unsigned int vertexShader, fragmentShader, shaderProgram;

    vertexShader = glCreateShader(GL_VERTEX_SHADER);
    glShaderSource(vertexShader, 1, &vertexShaderSourceImage, NULL);
    glCompileShader(vertexShader);

    fragmentShader = glCreateShader(GL_FRAGMENT_SHADER);
    glShaderSource(fragmentShader, 1, &fragmentShaderSourceImage, NULL);
    glCompileShader(fragmentShader);

    shaderProgram = glCreateProgram();
    glAttachShader(shaderProgram, vertexShader);
    glAttachShader(shaderProgram, fragmentShader);
    glLinkProgram(shaderProgram);

    // Check for shader compilation and linking errors here

    // Use shader program
    glUseProgram(shaderProgram);

    // Bind texture
    glActiveTexture(GL_TEXTURE0);
    glBindTexture(GL_TEXTURE_2D, texture);
    glUniform1i(glGetUniformLocation(shaderProgram, "ourTexture"), 0);

    // Draw quad
    glBindVertexArray(VAO);
    glDrawArrays(GL_TRIANGLES, 0, 4);

    // Cleanup
    glDeleteVertexArrays(1, &VAO);
    glDeleteBuffers(1, &VBO);
}

}  // namespace Vuelto
