namespace _18_TddFizzBuzzKata
{
    // FizzBuzz core logic
    public class FizzBuzzConverter
    {
        public string Convert(int value)
        {
            var result = string.Empty;

            if (value % 3 == 0)
            {
                result = "Fizz";
            }

            if (value % 5 == 0)
            {
                result += "Buzz";
            }

            if (result != string.Empty)
            {
                return result;
            }

            return value.ToString();
        }
    }
}
