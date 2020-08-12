import React from "react";
import useSwr from "swr";
import { makeStyles } from "@material-ui/core/styles";

import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";
import fetcher from "../api/fetcher";

const useStyles = makeStyles((theme) => ({
  table: {
    minWidth: 650,
    background: "#F2F2F2"
  },
  header: {
    background: "#FDFDFA",
  },
  headerLabel: {
    fontSize: theme.typography.pxToRem(16),
    fontWeight: theme.typography.fontWeightBold,
  }
}));

export default function ZoneRecords({ zone }) {
  const classes = useStyles();
  const { loading, data, error } = useSwr(`/api/records/${zone}`,fetcher)
  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead className={classes.header}>
          <TableRow>
            <TableCell className={classes.headerLabel}>Type</TableCell>
            <TableCell className={classes.headerLabel} align="left">Name</TableCell>
            <TableCell className={classes.headerLabel}>Content</TableCell>
            <TableCell className={classes.headerLabel} align="left">ttl</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {data && data.map((row) => (
            <TableRow key={row.id}>
              <TableCell component="th" scope="row">
                {row.type}
              </TableCell>
              <TableCell align="left">{`${row.name}.${row.zone_id}`}</TableCell>
              <TableCell>{row.content}</TableCell>
              <TableCell align="left">{row.ttl}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
