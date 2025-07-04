using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _13_Interfaces
{
    // Simple interface with one method
    interface IA
    {
        void foo();
    }

    // Another interface
    interface IB
    {
        void bar();
    }

    // Third interface
    interface IC
    {
        void qux();
    }

    // Interface inheritance
    interface ID : IA, IB { }

    // Class implements multiple interfaces
    class MyBase : IC, ID
    {
        public void bar()
        {
            // ...
        }

        public void foo()
        {
            // ,,,
        }

        public void qux()
        {
            // ;;;
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Interface method call through instance
            MyBase mb = new MyBase();
            mb.foo();
        }
    }
}
