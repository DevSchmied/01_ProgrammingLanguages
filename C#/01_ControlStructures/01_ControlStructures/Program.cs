using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _01_ControlStructures
{
    internal class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello C#!");


            Console.WriteLine("\nAufgabe 1");
            int[] mein_array = new int[7];

            Console.WriteLine(mein_array[3]);

            mein_array[2] = 11;

            Console.WriteLine("Index 2: " + mein_array[2]);
            Console.WriteLine("Index 3: " + mein_array[3]);
            Console.WriteLine("Index 4: " + mein_array[4]);


            // -----------------------------------
            Console.WriteLine("\nAufgabe 2");

            int[] arr = new int[7];

            for (int i = 0; i < 7; i++)
            {
                arr[i] = i * 3;
            }

            int index = 0;
            int counter = 0;
            while (index < 7)
            {
                if (arr[index] >= 5)
                {
                    counter++;
                }
                index++;
            }

            Console.WriteLine(counter);

            // -----------------------------------
            Console.WriteLine("\nAufgabe 3");

            int variable_1 = 10;

            int variable_2 = 2;

            while (variable_1 >= variable_2)
            {
                variable_1 += 1;
                variable_2 += 2;
            }

            Console.WriteLine(variable_1);
            Console.WriteLine(variable_2);

        }
    }
}
