"use client";

import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { useShowError } from "../../../context/Error";
import {
  Card,
  CardContent,
  Grid,
  Typography,
  Avatar,
  Box,
  Button,
  CircularProgress,
} from "@mui/material";
import { useRouter } from "next/navigation";

interface Team {
  Id: number;
  Name: string;
  City: string;
  StadiumId: number;
  Coach: string;
  Description: string;
  Founded: number;
  LogoUrl: string;
}

interface Stadium {
  Id: number;
  Name: string;
  Capacity: number;
  VipRows: number;
  SeatsPerRow: number;
}

interface MatchData {
  Id: number;
  HomeTeam: Team;
  AwayTeam: Team;
  Stadium: Stadium;
  Date: string;
  MainReferee: string;
  Lineman1: string;
  Lineman2: string;
}

export default function Page() {
  const { id } = useParams();
  const [data, setData] = useState<MatchData | null>(null);
  const showError = useShowError();
  const router = useRouter();

  useEffect(() => {
    if (!id) return;

    const fetchData = async () => {
      try {
        const response = await fetch(`http://localhost:8000/match/${id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const result: MatchData = await response.json();
        setData(result);
      } catch (err: any) {
        showError(err.message);
      }
    };

    fetchData();
  }, [id]);

  if (!data) {
    return (
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <CircularProgress />
      </Box>
    );
  }

  const handleReserveClick = () => {
    router.push(`/reserve/${data.Id}?stadiumId=${data.Stadium.Id}`);
  };

  return (
    <Box sx={{ padding: 3 }}>
      {/* Teams Section */}
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <Card sx={{ padding: 2, borderRadius: 2 }}>
            <CardContent>
              <Typography variant="h6" align="center" gutterBottom>
                {data.HomeTeam.Name}
              </Typography>
              <Avatar
                src={data.HomeTeam.LogoUrl}
                alt={data.HomeTeam.Name}
                sx={{ width: 100, height: 100, marginBottom: 2, mx: "auto" }}
              />
              <Typography variant="body1">
                <strong>City:</strong> {data.HomeTeam.City}
              </Typography>
              <Typography variant="body1">
                <strong>Coach:</strong> {data.HomeTeam.Coach}
              </Typography>
              <Typography variant="body1">
                <strong>Founded:</strong> {data.HomeTeam.Founded}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {data.HomeTeam.Description}
              </Typography>
            </CardContent>
          </Card>
        </Grid>

        <Grid item xs={12} md={6}>
          <Card sx={{ padding: 2, borderRadius: 2 }}>
            <CardContent>
              <Typography variant="h6" align="center" gutterBottom>
                {data.AwayTeam.Name}
              </Typography>
              <Avatar
                src={data.AwayTeam.LogoUrl}
                alt={data.AwayTeam.Name}
                sx={{ width: 100, height: 100, marginBottom: 2, mx: "auto" }}
              />
              <Typography variant="body1">
                <strong>City:</strong> {data.AwayTeam.City}
              </Typography>
              <Typography variant="body1">
                <strong>Coach:</strong> {data.AwayTeam.Coach}
              </Typography>
              <Typography variant="body1">
                <strong>Founded:</strong> {data.AwayTeam.Founded}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {data.AwayTeam.Description}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
      </Grid>

      {/* Stadium Section */}
      <Card sx={{ marginTop: 3, padding: 2, borderRadius: 2 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            Stadium Details
          </Typography>
          <Typography variant="body1">
            <strong>Name:</strong> {data.Stadium.Name}
          </Typography>
          <Typography variant="body1">
            <strong>Capacity:</strong> {data.Stadium.Capacity}
          </Typography>
          <Typography variant="body1">
            <strong>VIP Rows:</strong> {data.Stadium.VipRows}
          </Typography>
          <Typography variant="body1">
            <strong>Seats per Row:</strong> {data.Stadium.SeatsPerRow}
          </Typography>
        </CardContent>
      </Card>

      {/* Match Info Section */}
      <Card sx={{ marginTop: 3, padding: 2, borderRadius: 2 }}>
        <CardContent>
          <Typography variant="h6" gutterBottom>
            Match Information
          </Typography>
          <Typography variant="body1">
            <strong>Date:</strong> {new Date(data.Date).toLocaleDateString()}
          </Typography>
          <Typography variant="body1">
            <strong>Main Referee:</strong> {data.MainReferee}
          </Typography>
          <Typography variant="body1">
            <strong>Lineman 1:</strong> {data.Lineman1}
          </Typography>
          <Typography variant="body1">
            <strong>Lineman 2:</strong> {data.Lineman2}
          </Typography>
        </CardContent>
      </Card>

      {/* Reserve Button */}
      <Box sx={{ textAlign: "center", marginTop: 3 }}>
        <Button
          variant="contained"
          color="primary"
          size="large"
          onClick={handleReserveClick}
        >
          Reserve Seat
        </Button>
      </Box>
    </Box>
  );
}
