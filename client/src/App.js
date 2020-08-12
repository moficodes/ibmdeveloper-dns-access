import React from "react";
import Accordion from "@material-ui/core/Accordion";
import AccordionSummary from "@material-ui/core/AccordionSummary";
import AccordionDetails from "@material-ui/core/AccordionDetails";
import Typography from "@material-ui/core/Typography";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import { makeStyles } from "@material-ui/core/styles";
import useSwr from "swr";

import Table from "./components/Table";
import fetcher from "./api/fetcher";

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    marginBottom: '16px'
  },
  container: {
    margin: "32px",
  },
  heading: {
    fontSize: theme.typography.pxToRem(20),
    fontWeight: theme.typography.fontWeightBold,
  },
}));

export default function App() {
  const { loading, data, error } = useSwr("/zones", fetcher);
  const classes = useStyles();

  function renderZones(zones) {
    if (!zones) {
      return <></>;
    }
    return zones.map((zone) => (
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
            <Table zone={zone.name}/>
          </AccordionDetails>
        </Accordion>
      </div>
    ));
  }
  return (
    <div className={classes.container}>
      <div>{renderZones(data)}</div>
    </div>
  );
}
