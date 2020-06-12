import { SUITS, RANKS } from "react-playing-cards";

export const getSuit = (index: number): SUITS => {
  switch (index) {
    case 0:
      return SUITS.CLUBS;
    case 1:
      return SUITS.SPADES;
    case 2:
      return SUITS.HEARTS;
    case 3:
      return SUITS.DIAMONDS;
    default:
      return SUITS.CLUBS;
  }
};

export const getRank = (index: number): RANKS => {
  switch (index) {
    case 0:
      return RANKS.ACE;
    case 1:
      return RANKS.TWO;
    case 2:
      return RANKS.THREE;
    case 3:
      return RANKS.FOUR;
    case 4:
      return RANKS.FIVE;
    case 5:
      return RANKS.SIX;
    case 6:
      return RANKS.SEVEN;
    case 7:
      return RANKS.EIGHT;
    case 8:
      return RANKS.NINE;
    case 9:
      return RANKS.TEN;
    case 10:
      return RANKS.JACK;
    case 11:
      return RANKS.QUEEN;
    case 12:
      return RANKS.KING;
    default:
      return RANKS.ACE;
  }
};
