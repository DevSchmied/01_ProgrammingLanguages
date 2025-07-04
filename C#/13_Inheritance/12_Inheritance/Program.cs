using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _12_Inheritance
{
    // Base class with virtual method
    class Employee
    {
        public string Firstname { get; set; }
        public string Lastname { get; set; }
        public decimal Salary { get; set; }

        // Base income calculation
        virtual public decimal CalculateIncome()
        {
            return this.Salary;
        }
    }

    // Derived class with overridden method
    class Manager : Employee
    {
        public decimal Bonus { get; set; }

        // Manager income = salary + bonus
        public override decimal CalculateIncome()
        {
            return this.Salary + this.Bonus;
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Polymorphism: list of base class
            List<Employee> list = new List<Employee>();
            list.Add(new Employee() { Firstname = "Bob", Lastname = "Miller", Salary = 20000 });
            list.Add(new Manager() { Firstname = "John", Lastname = "Smith", Salary = 100000, Bonus = 50000 });
            list.Add(new Manager() { Firstname = "Sarah", Lastname = "Taylor", Salary = 80000, Bonus = 30000 });

            // Runtime method resolution
            foreach (var emp in list)
            {
                Console.WriteLine($"{emp.Firstname}, {emp.Lastname}, {emp.CalculateIncome()}$");
            }
        }
    }
}
