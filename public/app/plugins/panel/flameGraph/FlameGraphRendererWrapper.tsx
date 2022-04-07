import React from 'react';
import "@pyroscope/flamegraph/dist/index.css";
const flamegraph = require('@pyroscope/flamegraph');

const { FlamegraphRenderer } = flamegraph;

interface RenderWrapperProps {
  width: number;
  height: number;
  flamebearer: any;
}

export const FlameGraphRendererWrapper: React.FunctionComponent<RenderWrapperProps> = ({
  width,
  height,
  flamebearer,
}) => {
  return (
    <div style={{ width, height, overflowY: 'auto' }}>
      <FlamegraphRenderer
        flamebearer={flamebearer}
        ExportData={<div />}
        display="flamegraph"
        viewType="single"
        showToolbar={false}
      />
    </div>
  );
};
