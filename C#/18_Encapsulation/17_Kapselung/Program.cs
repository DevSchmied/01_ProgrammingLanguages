using System;

namespace _17_Encapsulation
{
    // Encapsulation and data hiding
    public class Vektor
    {
        int[] values;
        int length;

        // Constructor with length and default value
        public Vektor(int length, int value)
        {
            this.length = length;
            values = new int[length];
        }

        // Set single value
        public void SetValue(int index, int value)
        { values[index] = value; }

        // Get vector length
        public int GetLength()
        { return this.length; }

        // Initialize vector values 0..n
        public void Initialize()
        {
            for (int i = 0; i < values.Length; i++)
            { values[i] = i; }
        }

        // Multiply each value by a scalar
        public void MultiplyByScalar(int value)
        {
            for (int i = 0; i < values.Length; i++)
            { values[i] *= value; }
        }

        // Custom string representation
        public override string ToString()
        {
            string s = "";

            for (int i = 0; i < this.GetLength(); i++)
            {
                s += $"Element {i}: {values[i]}";
                if (i < values.Length - 1)
                    s += ",\n";
            }

            s += ".";

            return s;
        }
    }

    internal class Program
    {
        void DoIt()
        {
            // Vector usage example
            Vektor vektor = new Vektor(3, 2);
            vektor.Initialize();
            vektor.MultiplyByScalar(2);
            vektor.SetValue(2, 8);
            Console.WriteLine(vektor);
        }

        static void Main(string[] args)
        {
            new Program().DoIt();
        }
    }
}
