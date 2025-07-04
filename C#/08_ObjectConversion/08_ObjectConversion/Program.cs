using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _08_ObjectConversion
{
    // Employee class with properties and ToString()
    class Employee
    {
        public Employee(string firstname, string lastname)
        {
            this.firstname = firstname;
            this.lastname = lastname;
        }

        string firstname;
        public string Firstname
        {
            get { return this.firstname; }
            set { this.firstname = value; }
        }

        string lastname;
        public string Lastname
        {
            get { return this.lastname; }
            set { this.lastname = value; }
        }

        public override string ToString()
        {
            return this.firstname + " " + this.lastname;
        }
    }

    // Manager class with same structure as Employee
    class Manager
    {
        public Manager(string firstname, string lastname)
        {
            this.firstname = firstname;
            this.lastname = lastname;
        }

        string firstname;
        public string Firstname
        {
            get { return this.firstname; }
            set { this.firstname = value; }
        }

        string lastname;
        public string Lastname
        {
            get { return this.lastname; }
            set { this.lastname = value; }
        }

        public override string ToString()
        {
            return this.firstname + " " + this.lastname;
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            // Using non-generic ArrayList
            ArrayList al = new ArrayList();
            al.Add(new Employee("Bob", "Sturmer"));
            al.Add(new Manager("Klaus", "Franz"));

            // Type checking and casting with 'is' and cast
            for (int i = 0; i < al.Count; i++)
            {
                object obj = al[i];

                if (obj is Employee)
                {
                    Employee emp = (Employee)obj;
                    Console.WriteLine(emp);
                }
            }
        }
    }
}
