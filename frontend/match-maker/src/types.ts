export type User = {
  Username: string;
  Firstname: string;
  Lastname: string;
  Email: string;
  Gender: string;
  Role: string;
  Birthdate: string;
  City: string;
  Address: string;
  Approved: boolean;
};

export type Stadium = {
  Id?: number;
  Name: string;
  Capacity: number;
  VipRows: number;
  SeatsPerRow: number;
};

export type Team = {
  Id?: number;
  Name: string;
  City: string;
  StadiumId: number;
  Coach: string;
  Description: string;
  Founded: number;
  LogoUrl: string;
};

export type Match = {
  Id?: number;
  HomeTeam: Team;
  AwayTeam: Team;
  Stadium: Stadium;
  Date: string;
  MainReferee: string;
  Lineman1: string;
  Lineman2: string;
};

export type MatchRequest = {
  Id?: number;
  HomeTeamId: number;
  AwayTeamId: number;
  StadiumId: number;
  Date: string;
  MainReferee: string;
  Lineman1: string;
  Lineman2: string;
};
