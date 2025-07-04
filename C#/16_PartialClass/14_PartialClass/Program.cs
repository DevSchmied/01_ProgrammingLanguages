using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

// Demonstration of partial classes
namespace _14_PartialClass
{
    internal class Program
    {
        static void Main(string[] args)
        {
            // Call static methods from different parts of a partial class
            TestClass.foo();
            TestClass.bar();
        }
    }
}