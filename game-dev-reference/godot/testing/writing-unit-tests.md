```cs
namespace Tappy.Tests
{
    [TestSuite]
    public class PracticeTest
    {
        [TestCase]
        public void StringToLower() {
            AssertString("AbcD".ToLower()).IsEqual("abcd");
        }

        [TestCase]
        public void ProcessingStopsOnPlayerDeath() {
            Player player = AutoFree(new Player());
            player.SetProcess(true);
            player.OnDied();
            AssertThat(player.IsProcessing()).IsEqual(false);
        }
    }
}
```