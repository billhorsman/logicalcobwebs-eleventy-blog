export function toSentence(list) {
  if (list.length == 0) {
    return;
  } else if (list.length == 1) {
    return list[0];
  } else if (list.length == 2) {
    return list.join(" and ");
  } else {
    return list.slice(0, -1).join(", ") + " and " + list[list.length - 1];
  }
}
