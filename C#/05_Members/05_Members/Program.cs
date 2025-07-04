using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _05_Members
{
    public class Employee
    {
        // Private fields for first and last name
        string firstName;
        string lastName;

        // Public property for firstName with getter and setter
        public string FirstName
        {
            get { return this.firstName; }
            set { this.firstName = value; }
        }

        // Public property for lastName with getter and setter
        public string LastName
        {
            get { return this.lastName; }
            set { this.lastName = value; }
        }

        // Constructor to initialize firstName and lastName
        public Employee(string firstName, string lastName)
        {
            this.firstName = firstName;
            this.lastName = lastName;
        }

        // Method to print the employee's full name
        public void print()
        {
            Console.WriteLine(String.Format("First parameter: {0}, second parameter: {1}", this.firstName, this.lastName));
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Create an Employee object with first and last name
            Employee employee = new Employee("Mirko", "Matytschak");

            // Call print() method to display the names
            employee.print();

            // Print names using properties
            Console.WriteLine(employee.FirstName + " " + employee.LastName);

            int x;

            // Assign max value to int variable
            x = int.MaxValue;

            // Print max int value
            Console.WriteLine(x);
            /*
            // Checked block to detect overflow (commented out)
            checked
            {
                x = x + 1;
                Console.WriteLine(x);
            }
            */

            uint y;

            // Assign max value to uint variable
            y = uint.MaxValue;

            // Print max uint value
            Console.WriteLine(y);
            /*
            // Checked block to detect overflow (commented out)
            checked
            {
                y = y + 1;
                Console.WriteLine(y);
            }
            */
        }
    }
}
 