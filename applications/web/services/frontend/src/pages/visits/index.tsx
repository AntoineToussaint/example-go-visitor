import { useEffect } from "react";
import Layout from "../../components/layout";
import { useCodeflyContext } from "../providers/codefly.provider";
import VisitorCountChart from "./components/visitor-count-chart";


const Visits = () => {

  const { routing } = useCodeflyContext()

  const markVisit = () => {

    const visitUrl = routing("POST", { application: "backend", service: "server" }, "/visit")
    fetch(visitUrl,
      {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({})
      });
  }

  useEffect(() => {
    markVisit()
  }, [])


  return (
    <Layout>
      <div style={{ width: '60%', height: '400px', margin: '0 auto' }}>
        <VisitorCountChart />
      </div>
    </Layout>
  );
};


export default Visits;