import React from 'react';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import Header from './Header';
import MainFeaturedPost from './MainFeaturedPost';
import Footer from './Footer';
import Table from './Table';

const sections = [
  { title: 'Matches', url: '/dashboard' },
  { title: 'My profile', url: '/profile' },
  { title: 'Leaderboard', url: '/leaderboard' },
];

const mainFeaturedPost = {
  title: 'Leaderboard',
  description:
    "Top 100 players",
  image: 'https://source.unsplash.com/featured/?mathematics',
};

export default function LeaderboardPage() {

  return (
    <React.Fragment>
      <CssBaseline />
      <Container maxWidth="lg">
        <Header title="World Cup 2021 - Guess the score" sections={sections} />
        <main>
          <MainFeaturedPost post={mainFeaturedPost} />
          <Grid container spacing={4}>
            <Table/>
          </Grid>
        </main>
      </Container>
      <Footer title="PW 2021" description="Proiect implementat în cadrul materiei de Programare Web" />
     
     
    </React.Fragment>
  );
}