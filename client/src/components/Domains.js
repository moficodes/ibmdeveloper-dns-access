import React from "react";
import { makeStyles } from "@material-ui/core/styles";
import Accordion from "@material-ui/core/Accordion";
import AccordionSummary from "@material-ui/core/AccordionSummary";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import Typography from "@material-ui/core/Typography";
import AccordionDetails from "@material-ui/core/AccordionDetails";
import useSWR from "swr";
import fetcher from "../api/fetcher";
import ZoneRecord from "./ZoneRecords";

const useStyles = makeStyles((theme) => ({
  root: {
    margin: "16px",
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  heading: {
    fontSize: theme.typography.pxToRem(20),
    fontWeight: theme.typography.fontWeightBold,
  },
}));
export default function Domains() {
  const { loading, data, error } = useSWR("/api/zones", fetcher);
  const classes = useStyles();

  if (!data) {
    return <></>;
  }
  return data.map((zone) => (
    <div className={classes.root}>
      <Accordion>
        <AccordionSummary
          expandIcon={<ExpandMoreIcon />}
          aria-controls="panel1a-content"
          id="panel1a-header"
        >
          <Typography className={classes.heading}>{zone.name}</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <ZoneRecord zone={zone.name} />
        </AccordionDetails>
      </Accordion>
    </div>
  ));
}
