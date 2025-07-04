Console.WriteLine("Calculator");

// add()
Console.WriteLine("\nadd()-function");
int add(int value1, int value2)
{
  return value1 + value2;
}
Console.WriteLine(add(1, 2));
Console.WriteLine(add(-3, 2));
Console.WriteLine(add(0, 0));
Console.WriteLine(add(-5, -2));

// sub()
Console.WriteLine("\nsub()-function");
int sub(int value1, int value2)
{
  if (value1 > value2)
  {
    return value1 - value2;
  }
  else
  {
    return value2 - value1;
  }
}

Console.WriteLine(sub(5, 3));
Console.WriteLine(sub(5, -3));
Console.WriteLine(sub(-5, 3));
Console.WriteLine(sub(0, 0));

// multiply()
Console.WriteLine("\nmultiply()-function");
int multiply(int value1, int value2)
{
  bool isNegative = (value1 < 0 && value2 > 0) || (value1 > 0 && value2 < 0);

  if (value1 < 0)
  {
    value1 = -value1;
  }

  if (value2 < 0)
  {
    value2 = -value2;
  }

  int product = 0;
  for (int i = 0; i < value2; i++)
  {
    product += value1;
  }
  if (isNegative)
  {
    return -product;
  }
  return product;
}

Console.WriteLine(multiply(2, 2));
Console.WriteLine(multiply(3, 3));
Console.WriteLine(multiply(2, 4));
Console.WriteLine(multiply(-2, 2));
Console.WriteLine(multiply(-2, -2)); 
Console.WriteLine(multiply(-2, 6));


// divide()
Console.WriteLine("\ndivide()-function");
double divide(double value1, double value2)
{
  if (value2 == 0)
  {
    return 0;
  }

  return value1 / value2;
}

Console.WriteLine(divide(4, 2));
Console.WriteLine(divide(4, 0));
Console.WriteLine(divide(4.0, 2));
Console.WriteLine(divide(3, 2.0));
Console.WriteLine(divide(4, -2));
