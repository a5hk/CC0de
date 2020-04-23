from math import e

class ParentClass:
  def __init__(self):
    print("Parent class")

class ExampleClass(ParentClass):
  """Document"""

  classvar = "Hi"

  def __init__(self):
    super().__init__()
    self.num = 10

  def fibonacci(self, n):
    if n <= 0:
      return [0]

    sequence = [0, 1]
    while len(sequence) <= n:
      next_value = sequence[len(sequence) - 1] + sequence[len(sequence) - 2]
      sequence.append(next_value)

    return sequence


if __name__ == '__main__':
  obj = ExampleClass()
  print(obj.fibonacci(10))
  print(e)
