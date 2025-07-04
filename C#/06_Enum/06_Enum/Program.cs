using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _06_Enum
{
    // Enum definition
    enum Job
    {
        CheckStructure,
        GetNewData,
        RerenderPage
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Enum usage
            Job job = Job.CheckStructure;

            // Print enum name and value
            Console.WriteLine(job);
            Console.WriteLine((int)job);

            // Enum in if-statement
            if (job == Job.CheckStructure)
            {
                Console.WriteLine("CheckStructure");
            }

            // Enum in switch-statement
            switch (job)
            {
                case Job.CheckStructure:
                    Console.WriteLine("CheckStructure");
                    break;
                case Job.GetNewData:
                    Console.WriteLine("GetNewData");
                    break;
                case Job.RerenderPage:
                    Console.WriteLine("RerenderPage");
                    break;
                default:
                    Console.WriteLine("Nothing");
                    break;
            }
        }
    }
}
