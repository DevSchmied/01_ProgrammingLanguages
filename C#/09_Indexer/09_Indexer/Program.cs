using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _09_Indexer
{
    // Custom dictionary using indexer
    class MyDictionary
    {
        string[] keys = new string[10];
        int[] values = new int[10];

        string key;
        int value;

        // Indexer to get/set values by key
        public int this[string key]
        {
            get
            {
                for (int i = 0; i < keys.Length; i++)
                {
                    if (keys[i] == key)
                        return values[i];
                }
                throw new Exception("The key " + key + " was not found.");
            }
            set
            {
                for (int i = 0; i < keys.Length; i++)
                {
                    if (keys[i] == null)
                    {
                        keys[i] = key;
                        values[i] = value;
                        return;
                    }
                }
                throw new Exception("The dictionary has reached its capacity.");
            }
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Using custom indexer
            MyDictionary dict = new MyDictionary();

            dict["Hello"] = 1234;

            int hello = dict["Hello"];

            Console.WriteLine(hello);
        }
    }
}
