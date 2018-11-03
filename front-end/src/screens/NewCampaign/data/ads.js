import React from "react";
import { Header } from "semantic-ui-react";

export default [
  {
    text: "Tá com fome?",
    value: "hungry",
    content: (
      <Header
        as="h5"
        icon="money"
        content="Tá com fome?"
        subheader="Toma essa promoção!"
      />
    )
  },
  {
    text: "Sexta à noite...",
    value: "friday",
    content: (
      <Header
        as="h5"
        icon="moon"
        content="Sexta à noite..."
        subheader="Que tal um iFood?"
      />
    )
  },
  {
    text: "foods",
    value: 3,
    content: (
      <Header
        as="h5"
        icon="food"
        content="Peça foods!"
        subheader="15% até hoje"
      />
    )
  }
];
