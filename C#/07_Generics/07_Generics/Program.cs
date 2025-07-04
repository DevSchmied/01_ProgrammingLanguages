using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _07_Generics
{
    // Simple class with properties and ToString override
    class Employee
    {
        string firstname;
        string lastname;

        public string Firstname
        {
            get { return this.firstname; }
            set { this.firstname = value; }
        }

        public string Lastname
        {
            get { return this.lastname; }
            set { this.lastname = value; }
        }

        public override string ToString()
        {
            return "Employee: " + " " + this.firstname + " " + this.lastname;
        }

    }

    // Generic class with type constraint
    class Factory<T>
        where T : new()
    {
        public T createInstanz(string criteria)
        {
            // Create new instance of T
            return new T();
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Use of generic Factory
            Factory<Employee> f = new Factory<Employee>();

            // Create and initialize Employee object
            Employee e = f.createInstanz("...");
            e.Firstname = "Bob";
            e.Lastname = "Exampleman";

            // Print Employee info
            Console.WriteLine(e);
        }
    }
}
