import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import Header from './Header';
import MainFeaturedPost from './MainFeaturedPost';
import Paper from '@material-ui/core/Paper';
import Footer from './Footer';
import Typography from '@material-ui/core/Typography';
import Avatar from '../avatar.png'
import AuthService from "../../services/auth.service";


const useStyles = makeStyles((theme) => ({
  mainGrid: {
    marginTop: theme.spacing(3),
  },
  sidebarAboutBox: {
    padding: theme.spacing(2),
    backgroundColor: theme.palette.grey[100],
    textAlign: 'center',
  },
}));

const sections = [
  { title: 'Matches', url: '/dashboard' },
  { title: 'My profile', url: '/profile' },
  { title: 'Leaderboard', url: '/leaderboard' },
];

const mainFeaturedPost = {
  title: 'My profile',
  description:
    "This page is all about you",
  image: 'https://source.unsplash.com/random',
};

const user = AuthService.getCurrentUser()

export default function ProfilePage() {
  const classes = useStyles();

  return (
    <React.Fragment>
      <CssBaseline />
      <Container maxWidth="lg">
        <Header title="World Cup 2021 - Guess the score" sections={sections} />
        <main>
          <MainFeaturedPost post={mainFeaturedPost} />
          <Grid container spacing={3} justify="center" alignItems="center">
            <Grid item xs={12}>
                <Paper elevation={3} className={classes.sidebarAboutBox}>
                    <Typography gutterBottom>
                        <img src={Avatar} alt=""/>
                    </Typography>
                    <Typography variant="h6">
                        Email: {user.Email}
                    </Typography>
                    <Typography variant="h6">
                        Username: {user.Username}
                    </Typography>
                    <Typography variant="h6">
                        Score: {user.Score}
                    </Typography>
                    {/* <Typography variant="h6">
                        Predictions: 13
                    </Typography> */}
                </Paper>
            </Grid>
          </Grid>
        </main>
      </Container>
      <Footer title="PW 2021" description="Proiect implementat în cadrul materiei de Programare Web" />
     
     
    </React.Fragment>
  );
}