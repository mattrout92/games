import React from "react";
import AS from "../public/cards/AS.svg";
import TWOS from "../public/cards/2S.svg";
import THRS from "../public/cards/3S.svg";
import FOUS from "../public/cards/4S.svg";
import FIVS from "../public/cards/5S.svg";
import SIXS from "../public/cards/6S.svg";
import SEVS from "../public/cards/7S.svg";
import EIGS from "../public/cards/8S.svg";
import NINS from "../public/cards/9S.svg";
import TS from "../public/cards/TS.svg";
import JS from "../public/cards/JS.svg";
import QS from "../public/cards/QS.svg";
import KS from "../public/cards/KS.svg";
import AC from "../public/cards/AC.svg";
import TWOC from "../public/cards/2C.svg";
import THRC from "../public/cards/3C.svg";
import FOUC from "../public/cards/4C.svg";
import FIVC from "../public/cards/5C.svg";
import SIXC from "../public/cards/6C.svg";
import SEVC from "../public/cards/7C.svg";
import EIGC from "../public/cards/8C.svg";
import NINC from "../public/cards/9C.svg";
import TC from "../public/cards/TC.svg";
import JC from "../public/cards/JC.svg";
import QC from "../public/cards/QC.svg";
import KC from "../public/cards/KC.svg";
import AD from "../public/cards/AD.svg";
import TWOD from "../public/cards/2D.svg";
import THRD from "../public/cards/3D.svg";
import FOUD from "../public/cards/4D.svg";
import FIVD from "../public/cards/5D.svg";
import SIXD from "../public/cards/6D.svg";
import SEVD from "../public/cards/7D.svg";
import EIGD from "../public/cards/8D.svg";
import NIND from "../public/cards/9D.svg";
import TD from "../public/cards/TD.svg";
import JD from "../public/cards/JD.svg";
import QD from "../public/cards/QD.svg";
import KD from "../public/cards/KD.svg";
import AH from "../public/cards/AH.svg";
import TWOH from "../public/cards/2H.svg";
import THRH from "../public/cards/3H.svg";
import FOUH from "../public/cards/4H.svg";
import FIVH from "../public/cards/5H.svg";
import SIXH from "../public/cards/6H.svg";
import SEVH from "../public/cards/7H.svg";
import EIGH from "../public/cards/8H.svg";
import NINH from "../public/cards/9H.svg";
import TH from "../public/cards/TH.svg";
import JH from "../public/cards/JH.svg";
import QH from "../public/cards/QH.svg";
import KH from "../public/cards/KH.svg";
import BACK from "../public/cards/2B.svg";

export type Props = {
  descriptor: string;
  height: number;
  width: number;
};

export const Card = ({ descriptor, height, width }: Props): JSX.Element => {
  switch (descriptor) {
    case "AS":
      return <img src={AS} alt="card" height={height} width={width} />;
    case "2S":
      return <img src={TWOS} alt="card" height={height} width={width} />;
    case "3S":
      return <img src={THRS} alt="card" height={height} width={width} />;
    case "4S":
      return <img src={FOUS} alt="card" height={height} width={width} />;
    case "5S":
      return <img src={FIVS} alt="card" height={height} width={width} />;
    case "6S":
      return <img src={SIXS} alt="card" height={height} width={width} />;
    case "7S":
      return <img src={SEVS} alt="card" height={height} width={width} />;
    case "8S":
      return <img src={EIGS} alt="card" height={height} width={width} />;
    case "9S":
      return <img src={NINS} alt="card" height={height} width={width} />;
    case "10S":
      return <img src={TS} alt="card" height={height} width={width} />;
    case "JS":
      return <img src={JS} alt="card" height={height} width={width} />;
    case "QS":
      return <img src={QS} alt="card" height={height} width={width} />;
    case "KS":
      return <img src={KS} alt="card" height={height} width={width} />;
    case "AC":
      return <img src={AC} alt="card" height={height} width={width} />;
    case "2C":
      return <img src={TWOC} alt="card" height={height} width={width} />;
    case "3C":
      return <img src={THRC} alt="card" height={height} width={width} />;
    case "4C":
      return <img src={FOUC} alt="card" height={height} width={width} />;
    case "5C":
      return <img src={FIVC} alt="card" height={height} width={width} />;
    case "6C":
      return <img src={SIXC} alt="card" height={height} width={width} />;
    case "7C":
      return <img src={SEVC} alt="card" height={height} width={width} />;
    case "8C":
      return <img src={EIGC} alt="card" height={height} width={width} />;
    case "9C":
      return <img src={NINC} alt="card" height={height} width={width} />;
    case "10C":
      return <img src={TC} alt="card" height={height} width={width} />;
    case "JC":
      return <img src={JC} alt="card" height={height} width={width} />;
    case "QC":
      return <img src={QC} alt="card" height={height} width={width} />;
    case "KC":
      return <img src={KC} alt="card" height={height} width={width} />;
    case "AD":
      return <img src={AD} alt="card" height={height} width={width} />;
    case "2D":
      return <img src={TWOD} alt="card" height={height} width={width} />;
    case "3D":
      return <img src={THRD} alt="card" height={height} width={width} />;
    case "4D":
      return <img src={FOUD} alt="card" height={height} width={width} />;
    case "5D":
      return <img src={FIVD} alt="card" height={height} width={width} />;
    case "6D":
      return <img src={SIXD} alt="card" height={height} width={width} />;
    case "7D":
      return <img src={SEVD} alt="card" height={height} width={width} />;
    case "8D":
      return <img src={EIGD} alt="card" height={height} width={width} />;
    case "9D":
      return <img src={NIND} alt="card" height={height} width={width} />;
    case "10D":
      return <img src={TD} alt="card" height={height} width={width} />;
    case "JD":
      return <img src={JD} alt="card" height={height} width={width} />;
    case "QD":
      return <img src={QD} alt="card" height={height} width={width} />;
    case "KD":
      return <img src={KD} alt="card" height={height} width={width} />;
    case "AH":
      return <img src={AH} alt="card" height={height} width={width} />;
    case "2H":
      return <img src={TWOH} alt="card" height={height} width={width} />;
    case "3H":
      return <img src={THRH} alt="card" height={height} width={width} />;
    case "4H":
      return <img src={FOUH} alt="card" height={height} width={width} />;
    case "5H":
      return <img src={FIVH} alt="card" height={height} width={width} />;
    case "6H":
      return <img src={SIXH} alt="card" height={height} width={width} />;
    case "7H":
      return <img src={SEVH} alt="card" height={height} width={width} />;
    case "8H":
      return <img src={EIGH} alt="card" height={height} width={width} />;
    case "9H":
      return <img src={NINH} alt="card" height={height} width={width} />;
    case "10H":
      return <img src={TH} alt="card" height={height} width={width} />;
    case "JH":
      return <img src={JH} alt="card" height={height} width={width} />;
    case "QH":
      return <img src={QH} alt="card" height={height} width={width} />;
    case "KH":
      return <img src={KH} alt="card" height={height} width={width} />;
    case "back":
      return <img src={BACK} alt="card" height={height} width={width} />;
    default:
      return <img src={AS} alt="card" height={height} width={width} />;
  }
};
