import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import Header from './Header';
import MainFeaturedPost from './MainFeaturedPost';
import FeaturedPost from './FeaturedPost';
import Footer from './Footer';
import BusinessService from "../../services/business.service";
import Button from '@material-ui/core/Button';
import AuthService from "../../services/auth.service";

const useStyles = makeStyles((theme) => ({
  mainGrid: {
    marginTop: theme.spacing(3),
  },
  pad: {
    paddingBottom: 10
  }
}));

const sections = [
  { title: 'Matches', url: '/dashboard' },
  { title: 'My profile', url: '/profile' },
  { title: 'Leaderboard', url: '/leaderboard' },
];

const mainFeaturedPost = {
  title: 'Can You Guess the World Cup Final Score ?',
  description:
    "Football Quiz",
  image: 'https://source.unsplash.com/featured/?soccer',
};

export default function Dashboard() {
  const classes = useStyles();
  let [featuredPosts, setFeaturedPosts] = useState([])

  useEffect(async () => {
    let results = await BusinessService.allMatches()
    setFeaturedPosts(results)
  },[]);

  return (
    <React.Fragment>
      <CssBaseline />
      <Container maxWidth="lg">
        <Header title="World Cup 2021 - Guess the score" sections={sections} />
        <main>
          <MainFeaturedPost post={mainFeaturedPost} />
          <Grid container justify="flex-end" className={classes.pad}>
          {
            AuthService.canEdit() &&
            <Button variant="contained" color="primary" > Add match </Button>
          }
          </Grid>

          <Grid container spacing={4}>
            {featuredPosts.map((post) => (
              <FeaturedPost key={post._id} post={post} />
            ))}
          </Grid>
        </main>
      </Container>
      <Footer title="PW 2021" description="Proiect implementat în cadrul materiei de Programare Web" />
     
     
    </React.Fragment>
  );
}