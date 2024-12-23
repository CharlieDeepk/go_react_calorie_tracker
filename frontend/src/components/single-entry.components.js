import React from "react";
import "bootstrap/dist/css/bootstrap.css";
import { Button, Card, Row, Col } from "react-bootstrap";

export default function Entry({
  entryData,
  setChangeIngredient,
  deleteSingleEntry,
  setChangeEntry,
}) {
  const handleChangeIngredient = () => {
    setChangeIngredient({
      change: true,
      id: entryData._id,
    });
  };
  const handleChangeEntry = () => {
    setChangeEntry({
      change: true,
      id: entryData._id,
    });
  };
  return (
    <Card>
      <Row>
        <Col>Dish{entryData !== undefined && entryData.dish}</Col>
        <Col>Ingredients{entryData !== undefined && entryData.ingredients}</Col>
        <Col>Calories{entryData !== undefined && entryData.calories}</Col>
        <Col>Fats{entryData !== undefined && entryData.fats}</Col>
        <Col>
          <Button onClick={() => deleteSingleEntry(entryData._id)}>
            Delete
          </Button>
        </Col>
        <Col>
          <Button onClick={() => handleChangeIngredient()}>
            Change Ingredient
          </Button>
        </Col>
        <Col>
          <Button onClick={() => handleChangeEntry()}>Change Entry</Button>
        </Col>
      </Row>
    </Card>
  );
}
