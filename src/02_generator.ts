export function* generator() {
  console.log("Strat...🕐");
  let value = 1;
  while (value < 4) yield value++;
  console.log("Finished!");
}
