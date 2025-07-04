using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _03_GuessNumber
{
    internal class Program
    {
        static void Main(string[] args)
        {
            int theNumber = new Random().Next(20);

            Console.WriteLine("Let's play the game 'Guess the Number'!");
            Console.WriteLine("I'm thinking of a number between 0 and 20.");

            int guessNum = 21;
            int counter = 1;
            while (guessNum != -1 || guessNum == theNumber)
            {
                Console.WriteLine("Try to guess my number, or enter -1 to quit the game.");
                guessNum = Convert.ToInt32(Console.ReadLine());
                if (guessNum == -1)
                {
                    Console.WriteLine("Too bad. The number was " + theNumber + ". Bye!");
                    return;
                }
                if (guessNum > theNumber)
                {
                    Console.WriteLine("The number is smaller!");
                    counter++;
                }
                else if (guessNum < theNumber)
                {
                    Console.WriteLine("The number is bigger!");
                    counter++;
                }
                else
                {
                    Console.WriteLine("Correct! The number is " + guessNum + ".\nYou needed " + counter + " attempts.");
                    return;
                }
            }
        }
    }
}
