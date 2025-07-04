using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _11_Exceptions
{
    internal class Program
    {
        static void Main(string[] args)
        {
            int a = 5;
            int b = 0;
            int result;

            try
            {
                // Division by zero - triggers exception
                result = a / b;
            }
            catch (Exception ex)
            {
                // Exception handling and message output
                Console.WriteLine("Exception message: \"" + ex.Message + "\"");
            }
            finally
            {
                // Finally block always runs
                Console.WriteLine("Always executed");
            }

            // Code after exception
            Console.WriteLine("Program continued after exception.");
        }
    }
}
