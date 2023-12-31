import React from "react";
// `Appointment` is not the default export, which is intentional
import { Appointment, AppointmentsDayView } from "../src/AppointmentsDayView";
import { render, click, describeReactComponent } from "./reactTestExtensions";

describeReactComponent("Appointment", () => {
  it("renders the customer first name", () => {
    const customer = {
      firstName: "Ashley",
    };
    render(<Appointment customer={customer} />);
    expect(document.body).toContainText("Ashley");
  });

  it("renders another customer first name", () => {
    const customer = {
      firstName: "Jordan",
    };
    render(<Appointment customer={customer} />);
    expect(document.body).toContainText("Jordan");
  });
});

describeReactComponent("AppointmentsDayView", () => {
  //#region setup
  const today = new Date();
  const twoAppointments = [
    {
      startsAt: today.setHours(12, 0),
      customer: { firstName: "Ashley" },
    },
    {
      startsAt: today.setHours(13, 0),
      customer: { firstName: "Jordan" },
    },
  ];
  //#endregion
  it("renders a div with the right id", () => {
    render(<AppointmentsDayView appointments={[]} />);
    expect(document.querySelector("div#appointmentsDayView")).not.toBeNull();
  });
  it("renders on ol element to display appointments", () => {
    render(<AppointmentsDayView appointments={[]} />);
    const listElement = document.querySelector("ol");
    expect(listElement).not.toBeNull();
  });
  it("renders an li for each appointment", () => {
    render(<AppointmentsDayView appointments={twoAppointments} />);
    const listOfChildren = document.querySelectorAll("ol > li");
    expect(listOfChildren).toHaveLength(2);
  });
  it("renders the time of each appointment", () => {
    render(<AppointmentsDayView appointments={twoAppointments} />);
    const listOfChildren = document.querySelectorAll("li");
    expect(listOfChildren[0]).toContainText("12:00");
    expect(listOfChildren[1]).toContainText("13:00");
  });
  it("initially shows a message saying there are no appointments today", () => {
    render(<AppointmentsDayView appointments={[]} />);
    expect(document.body).toContainText(
      "There are no appointments scheduled for today"
    );
  });
  it("selects the first appointment by default", () => {
    render(<AppointmentsDayView appointments={twoAppointments} />);
    expect(document.body).toContainText("Ashley");
  });
  it("has a button element in each li", () => {
    render(<AppointmentsDayView appointments={twoAppointments} />);
    const buttons = document.querySelectorAll("li > button");
    expect(buttons).toHaveLength(2);
    expect(buttons[0].type).toEqual("button");
  });
  it("renders another appointment when selected", () => {
    render(<AppointmentsDayView appointments={twoAppointments} />);
    const button = document.querySelectorAll("button")[1];
    click(button);
    expect(document.body).toContainText("Jordan");
  });
});
