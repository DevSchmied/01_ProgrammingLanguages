using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _10_Loops
{
    internal class Program
    {
        static void Main(string[] args)
        {
            // 2D array
            int[,] arr = {
                { 0,  1,  2,  3,  4,  5,  6,  7,  8,  9 },
                { 10, 11, 12, 13, 14, 15, 16, 17, 18, 19 }
            };

            Console.WriteLine("\n--------------FOR--------------\n");
            // Nested for loop
            for (int i = 0; i < arr.GetLength(0); i++)
            {
                for (int j = 0; j < arr.GetLength(1); j++)
                {
                    Console.WriteLine(string.Format("arr[{0}][{1}]: {2}", i, j, arr[i, j]));
                }
            }

            // Jagged array
            int[][] arr2 = {
                new int [] { 0,  1,  2,  3,  4,  5,  6,  7,  8,  9 },
                new int[]  { 10, 11, 12, 13, 14, 15, 16, 17, 18 }
            };

            // Loop through jagged array
            Console.WriteLine("\n----------------------------\n");
            for (int i = 0; i < arr2.Length; i++)
            {
                for (int j = 0; j < arr2[i].Length; j++)
                {
                    Console.WriteLine(string.Format("arr2[{0}][{1}]: {2}", i, j, arr2[i][j]));
                }
            }

            Console.WriteLine("\n--------------WHILE--------------\n");
            // While loop
            int k = 0;
            int l = 0;

            while (k < arr.GetLength(0))
            {
                while (l < arr.GetLength(1))
                {
                    Console.WriteLine(string.Format("arr[{0}][{1}]: {2}", k, l, arr[k, l]));
                    l++;
                }
                l = 0;
                k++;
            }

            Console.WriteLine("\n--------------DO-WHILE--------------\n");
            // Do-while loop with continue and break
            k = 0;
            l = 0;

            do
            {
                if (k % 2 == 1)
                {
                    k++;
                    continue;
                }

                l = 0;

                do
                {
                    if (l > 6)
                    {
                        break;
                    }

                    Console.WriteLine(string.Format("arr[{0}][{1}]: {2}", k, l, arr[k, l]));
                    l++;
                } while (l < arr.GetLength(1));

                k++;
            } while (k < arr.GetLength(0));

            Console.WriteLine("\n--------------FOREACH--------------\n");
            // Foreach loop on jagged array
            foreach (int[] i in arr2)
            {
                foreach (int j in i)
                {
                    Console.WriteLine(string.Format("arr2[][]: {2}", i, j, j));
                }
            }
        }
    }
}
