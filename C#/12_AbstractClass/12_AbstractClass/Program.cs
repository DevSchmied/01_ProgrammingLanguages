using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _12_AbstractClass
{
    // Abstract base class
    abstract class Person
    {
        // Auto-properties for personal data
        string firstname;
        public string Firstname { get; set; }

        string lastname;
        public string Lastname { get; set; }

        string street;
        public string Street { get; set; }

        string city;
        public string City { get; set; }
    }

    // Derived class without additional members
    class Customer : Person
    {
    }

    // Derived class with method
    class Employee : Person
    {
        decimal salary;

        // Method of derived class
        public void Work()
        {
            Console.WriteLine("I do my work!");
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Abstract class cannot be instantiated
            // Person person = new Person();

            // Create instances of derived classes
            Customer customer = new Customer() { Firstname = "Tom", Lastname = "Johnson", City = "New York", Street = "Avenue Park" };
            Employee employee = new Employee() { Firstname = "Sara", Lastname = "Taylor", City = "Chicago", Street = "Rosen Street" };
        }
    }
}
