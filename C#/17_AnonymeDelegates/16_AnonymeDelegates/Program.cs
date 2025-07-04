using System;
using System.Collections.Generic;

namespace _16_AnonymeDelegates
{
    internal class Program
    {
        static void Main(string[] args)
        {
            // List with integers
            List<int> list = new List<int>() { 5, 6, 2, 90, 1, 34 };

            Console.WriteLine("\n--------- Descending sort with anonymous delegate ---------\n");

            // Anonymous delegate - descending sort
            list.Sort(delegate (int i1, int i2)
            {
                return i2 - i1;
            });

            foreach (int i in list)
            {
                Console.WriteLine(i);
            }

            Console.WriteLine("\n--------- Ascending sort with lambda expression ---------\n");

            // Lambda expression - ascending sort
            list.Sort((i1, i2) =>
            {
                return i1 - i2;
            });

            foreach (int i in list)
            {
                Console.WriteLine(i);
            }
        }
    }
}
