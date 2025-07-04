using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _04_Bubblesort
{
    internal class Program
    {
        static void Main(string[] args)
        {
            void BubbleSort(int[] toSort)
            {
                if (toSort == null || toSort.Length == 0)
                {
                    Console.WriteLine("The array must not be null and must contain at least one element!");
                    return;
                }
                int temp = 0;
                for (int i = 0; i < toSort.Length; i++)
                {
                    for (int j = 0; j < toSort.Length - 1 - i; j++)
                    {
                        if (toSort[j] > toSort[j + 1])
                        {
                            temp = toSort[j];
                            toSort[j] = toSort[j + 1];
                            toSort[j + 1] = temp;
                        }
                    }
                }


            }


            int[] my_array = { 22, 1, 21, 11, 5 };


            BubbleSort(my_array);

            for (int i = 0; i < my_array.Length; i++)
            {
                Console.WriteLine(my_array[i]);
            }
        }
    }
}
