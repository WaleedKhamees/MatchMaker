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
  Id?: string;
  HomeTeam: Team;
  AwayTeam: Team;
  StadiumId: number;
  Date: string;
  MainReferee: string;
  LineMan1: string;
  LineMan2: string;
};
