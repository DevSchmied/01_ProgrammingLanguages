using System;
using System.Collections.Generic;

namespace _15_Delegates
{
    class Program
    {
        // Custom delegate declaration
        delegate bool StringTester(string s);

        static void Main(string[] args)
        {
            // Run logic
            new Program().DoIt();

        }

        void DoIt()
        {
            // Using delegate with whitespace check
            StringTester st1 = new StringTester(ContainsWhiteSpace);
            Console.WriteLine(st1("Ab af")); // true
            Console.WriteLine(st1("abc"));   // false

            // Using delegate with digit check
            StringTester st2 = new StringTester(ContainsDigit);
            Console.WriteLine(st2("Ab2af")); // true
            Console.WriteLine(st2("abc"));   // false
        }

        // Checks if string contains whitespace
        bool ContainsWhiteSpace(string s)
        {
            for (int i = 0; i < s.Length; i++)
            {
                if (Char.IsWhiteSpace(s[i]))
                {
                    return true;
                }
            }
            return false;
        }

        // Checks if string contains a digit
        bool ContainsDigit(string s)
        {
            for (int i = 0; i < s.Length; i++)
            {
                if (Char.IsDigit(s[i]))
                {
                    return true;
                }
            }
            return false;
        }
    }
}
